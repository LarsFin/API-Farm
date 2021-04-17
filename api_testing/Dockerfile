FROM postman/newman:latest

COPY ./postman/* /etc/newman/

ENTRYPOINT ["newman"]

#  "--folder Tests", "-e Development.postman_environment.json", "--env-var host=ruby_sinatra"
CMD ["run", "API_farm.postman_collection.json"]