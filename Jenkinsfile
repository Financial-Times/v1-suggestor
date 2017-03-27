node('docker') {
  stage 'checkout'
  checkout scm
  
  stage 'build-image'
  docker.build("coco/v1-suggestor:pipeline01", ".") 
}