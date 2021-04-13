# frozen_string_literal: true

require 'sinatra'

require_relative 'lib/controller/controller'
require_relative 'lib/models/video_game'
require_relative 'lib/services/video_games'
require_relative 'lib/storage/in_memory'

# environments; DEV, TEST, PROD
environment = ENV['RUBY_SINATRA_ENV'] || 'DEV'
environment.upcase!

storage = InMemory.new
service = VideoGames.new storage, VideoGame
controller = Controller.new service

get '/' do
    'Hello World!'
end

if environment == 'TEST'
    get '/api_tests/setup' do
        'TBD'
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

def get_params
    params.each { |key, value| request.params[key] = value }
end

