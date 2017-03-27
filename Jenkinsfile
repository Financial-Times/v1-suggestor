node {
    agent any
    stages {
        stage('build-image') {
            steps {
                sh 'docker build -t coco/v1-suggestor:pipeline01 .'
            }
        }
    }
}