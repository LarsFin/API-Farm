/*
There are three scenarios the pipeline runs against;
1) The target branch is master and current branch is lang/framework; linting, src code testing & api testing should be run
2) The target branch is lang/framework; linting & src code testing should be run
3) The target branch is neither of the above; no steps should be run
*/

// Append new lang/framework branch names when ready for pipeline builds
def langFrameworks = [
    'ruby/sinatra',
    'ruby/sinatra_jenkins_test_2'
]

def isIntoMaster = false
def isIntoLangFramework = false
def buildPath = ''
def buildService = ''
def apiTestPath = ''

if (env.CHANGE_TARGET == 'master' && langFrameworks.contains(env.CHANGE_BRANCH)) {
    isIntoMaster = true
    buildPath = 'ruby/sinatra' // "${env.CHANGE_BRANCH}"
    buildService = 'ruby_sinatra' // buildPath.replace('/', '_')
    apiTestPath = "api_testing"
} else if (langFrameworks.contains(env.CHANGE_TARGET)) {
    isIntoLangFramework = true
    buildPath = "${env.CHANGE_TARGET}"
}

def formatError(e) {
    "```\nMessage: ${e.message.substring(0, Math.min(255, e.message.length()))}\nStack Trace: ${e.stackTrace.substring(0, Math.min(255, e.stackTrace.length()))}\n```"
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
                script {
                    try {
                        echo "Running build script; ${buildPath}/scripts/build_img.sh"
                        dir(buildPath) {
                            sh 'chmod 700 -R ./scripts'
                            sh './scripts/build_img.sh'
                        }
                    } catch (e) {
                        pullRequest.comment("BUILD FAILED ❌. SEE ERROR DETAILS BELOW:\n${formatError(e)}")
                        throw e
                    }
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
                script {
                    try {
                        echo "Running lint script; ${buildPath}/scripts/run_img.sh lint"
                        dir(buildPath) {
                            sh './scripts/run_img.sh lint'
                        }
                    } catch (e) {
                        pullRequest.comment("LINTING FAILED ❌. SEE ERROR DETAILS BELOW:\n${formatError(e)}")
                        throw e
                    }
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
                script {
                    try {
                        echo "Running src test script; ${buildPath}/scripts/run_img.sh test"
                        dir(buildPath) {
                            sh './scripts/run_img.sh test'
                        }   
                    } catch (e) {
                        pullRequest.comment("SRC TESTING FAILED ❌. SEE ERROR DETAILS BELOW:\n${formatError(e)}")
                        throw e
                    }
                }
            }
        }

        stage('Health Check') {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                script {
                     try {
                        echo "Running api script; ${buildPath}/scripts/run_img.sh"
                        dir(buildPath) {
                            sh './scripts/run_img.sh'
                        }
                        sh 'curl -f http://localhost:8080/ping'
                    } catch (e) {
                        echo "${e.dump()}"
                        pullRequest.comment("HEALTH CHECK FAILED ❌. SEE ERROR DETAILS BELOW:\n${formatError(e)}")
                        throw e
                    }
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
                script {
                    try {
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
                    } catch (e) {
                        pullRequest.comment("API TESTS FAILED ❌. SEE ERROR DETAILS BELOW:\n${formatError(e)}")
                        throw e
                    }
                }
            }
        }

        stage('Finish') {
            steps {
                script {
                    pullRequest.comment("BUILD SUCCESSFUL ✔️")
                }
            }
        }
    }
}