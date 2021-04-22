pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'printenv'
                echo 'Building...'
            }
        }

        stage('Test') {
            steps {
                echo 'Testing...'
            }
        }

        stage('Prod Check') {
            steps {
                when {
                    env.CHANGE_TARGE == "master"
                }
                stage {
                    echo "Running API Tests..."
                }

                echo 'Complete!'
            }
        }
    }
}