pipeline {
    agent any

    stages {
        stage('Descargar Código') {
            steps {
                checkout scm
            }
        }

        stage('Construir Imagen Docker') {
            steps {
                sh 'docker build -t mi-web-uoc .'
            }
        }

        stage('Desplegar Web') {
            steps {
                sh 'docker stop mi-web-uoc || true'
                sh 'docker rm mi-web-uoc || true'
                sh 'docker run -d --name mi-web-uoc -p 8081:8080 mi-web-uoc'
            }
        }
    }
}
