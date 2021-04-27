/*
There are three scenarios the pipeline runs against;
1) The target branch is master and current branch is lang/framework; linting, src code testing & api testing should be run
2) The target branch is lang/framework; linting & src code testing should be run
3) The target branch is neither of the above; no steps should be run
*/

// Append new lang/framework branch names when ready for pipeline builds
def langFrameworks = [
    'ruby/sinatra'
]

def isIntoMaster = false
def isIntoLangFramework = false
def buildPath = ''
def buildService = ''
def apiTestPath = ''

if (env.CHANGE_TARGET == 'master' && langFrameworks.contains(env.CHANGE_BRANCH)) {
    isIntoMaster = true
    buildPath = "${env.CHANGE_BRANCH}"
    buildService = buildPath.replace('/', '_')
    apiTestPath = "api_testing"
} else if (langFrameworks.contains(env.CHANGE_TARGET)) {
    isIntoLangFramework = true
    buildPath = "${env.CHANGE_TARGET}"
}

pipeline {
    agent any

    stages {
        stage('Build') {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                echo "Running build script; ${buildPath}/scripts/build_img.sh"
                dir(buildPath) {
                    sh 'chmod 700 -R ./scripts'
                    sh './scripts/build_img.sh'
                }
            }
        }

        stage('Linting') {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                echo "Running lint script; ${buildPath}/scripts/run_img.sh lint"
                dir(buildPath) {
                    sh './scripts/run_img.sh lint'
                }
            }
        }

        stage('Src Testing') {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                echo "Running src test script; ${buildPath}/scripts/run_img.sh test"
                dir(buildPath) {
                    sh './scripts/run_img.sh test'
                }
            }
        }

        stage('API Testing') {
            when {
                expression {
                    return isIntoMaster
                }
            }
            steps {
                echo "Running api script; ${buildPath}/scripts/run_img.sh"
                dir(buildPath) {
                    sh './scripts/run_img.sh'
                }
                echo "Running expectations api build script; ${apiTestPath}/expectations_api/scripts/build_img.sh"
                dir("${apiTestPath}/expectations_api") {
                    sh 'chmod 700 -R ./scripts'
                    sh './scripts/build_img.sh'
                    echo "Running expectations api script; ${apiTestPath}/expectations_api/scripts/run_img.sh"
                    sh './scripts/run_img.sh'
                }
                echo "Running api tests!"
                dir(apiTestPath) {
                    sh 'chmod 700 ./run.sh'
                    sh "./run.sh ${buildService}"
                }
            }
        }

        stage('Finish') {
            steps {
                echo "Complete!"
            }
        }
    }
}