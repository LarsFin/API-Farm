# frozen_string_literal: true

require 'sinatra'

require_relative 'lib/controller/controller'
require_relative 'lib/models/video_game'
require_relative 'lib/services/video_games'
require_relative 'lib/storage/in_memory'
require_relative 'lib/utils/config'

# environments; DEV, TEST, PROD
environment = (ENV['API_ENV'] || 'DEV').upcase

config = Config.new environment

# configure host and port
set :bind, config.settings['bind']
set :port, config.settings['port']

storage = InMemory.new
service = VideoGames.new storage, VideoGame
controller = Controller.new service

get '/ping' do
    'pong'
end

if %w[TEST DEV].include? environment
    require_relative 'lib/controller/testing_controller'
    require_relative 'lib/services/data_loader'

    data_loader = DataLoader.new 'data.json', VideoGame
    testing_controller = TestingController.new data_loader, storage

    get '/api_tests/setup' do
        testing_controller.setup

        'Successfully loaded data.'
    end
end

get '/video_games' do
    controller.get_video_games
end

post '/video_games' do
    controller.add_video_game request
end

get '/video_games/:id' do
    get_params
    controller.get_video_game request
end

put '/video_games/:id' do
    get_params
    controller.update_video_game request
end

delete '/video_games/:id' do
    get_params
    controller.delete_video_game request
end

def get_params
    params.each { |key, value| request.params[key] = value }
end
