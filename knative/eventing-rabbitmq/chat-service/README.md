## Important commands
docker run --rm -it -p 15672:15672 -p 5672:5672 chat-service
Remember to forward ports

kubectl port-forward svc/rabbitmq-service 15672:80

kubectl port-forward svc/chat-service 32080:8080
