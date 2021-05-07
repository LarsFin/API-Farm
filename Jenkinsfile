/*
There are three scenarios the pipeline runs against;
1) The target branch is master and current branch is lang/framework; linting, src code testing & api testing should be run
2) The target branch is lang/framework; linting & src code testing should be run
3) The target branch is neither of the above; no steps should be run
*/

// Append new lang/framework branch names when ready for pipeline builds
def langFrameworks = [
    "js/express",
    "ruby/sinatra"
]

def isIntoMaster = false
def isIntoLangFramework = false
def buildPath = ""
def buildService = ""
def apiTestPath = ""

if (env.CHANGE_TARGET == "master" && langFrameworks.contains(env.CHANGE_BRANCH)) {
    isIntoMaster = true
    buildPath = "${env.CHANGE_BRANCH}"
    buildService = buildPath.replace('/', '_')
    apiTestPath = "api_testing"
} else if (langFrameworks.contains(env.CHANGE_TARGET)) {
    isIntoLangFramework = true
    buildPath = "${env.CHANGE_TARGET}"
}

def formatMessage(msg) {
    "```\n${msg.substring(0, Math.min(255, msg.length()))}\n```"
}

pipeline {
    agent any

    stages {
        stage("Build") {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                script {
                    try {
                        dir(buildPath) {
                            sh "chmod 700 -R ./scripts"
                            sh "find . -name \"*.json\" -exec chmod 600 {} \\;"
                        }

                        echo "Running build script; ./scripts/build.sh ${buildPath}"
                        sh "chmod 700 -R ./scripts"
                        sh "./scripts/build.sh ${buildPath}"
                    } catch (e) {
                        pullRequest.comment("BUILD FAILED ❌. SEE ERROR MESSAGE BELOW:\n${formatMessage(e.message)}")
                        throw e
                    }
                }
            }
        }

        stage("Linting") {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                script {
                    try {
                        echo "Running lint script; ./scripts/run.sh ${buildPath} lint"
                        sh "./scripts/run.sh ${buildPath} lint"
                    } catch (e) {
                        pullRequest.comment("LINTING FAILED ❌. SEE ERROR MESSAGE BELOW:\n${formatMessage(e.message)}")
                        throw e
                    }
                }
            }
        }

        stage("Src Testing") {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                script {
                    try {
                        echo "Running src test script; ./scripts/run.sh ${buildPath} test"
                        sh "./scripts/run.sh ${buildPath} test"
                    } catch (e) {
                        pullRequest.comment("SRC TESTING FAILED ❌. SEE ERROR MESSAGE BELOW:\n${formatMessage(e.message)}")
                        throw e
                    }
                }
            }
        }

        stage("Health Check") {
            when {
                expression {
                    return isIntoMaster || isIntoLangFramework
                }
            }
            steps {
                script {
                     try {
                        echo "Running api script; ./scripts/run.sh ${buildPath}"
                        sh "./scripts/run.sh ${buildPath}"
                        sh "sleep 2"
                        sh "curl -f http://localhost:8080/ping"
                    } catch (e) {
                        pullRequest.comment("HEALTH CHECK FAILED ❌. SEE ERROR MESSAGE BELOW:\n${formatMessage(e.message)}")
                        throw e
                    }
                }
            }
        }

        stage("API Testing") {
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
                            sh "chmod 700 -R ./scripts"
                            sh "./scripts/build_img.sh"
                            echo "Running expectations api script; ${apiTestPath}/expectations_api/scripts/run_img.sh"
                            sh "./scripts/run_img.sh"
                        }
                        echo "Running api tests!"
                        dir(apiTestPath) {
                            sh "chmod 700 ./run.sh"
                            sh "./run.sh ${buildService}"
                        }
                    } catch (e) {
                        pullRequest.comment("API TESTS FAILED ❌. SEE ERROR MESSAGE BELOW:\n${formatMessage(e.message)}")
                        throw e
                    }
                }
            }
        }

        stage("Finish") {
            steps {
                script {
                    pullRequest.comment("BUILD SUCCESSFUL ✔️")
                }
            }
        }
    }
    post {
        always {
            script {
                if (isIntoMaster || isIntoLangFramework) {
                    echo "Running clean up script; ./scripts/clean.sh ${buildPath}"
                    sh "./scripts/clean.sh ${buildPath}"
                }
            }
        }
    }
}