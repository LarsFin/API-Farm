# frozen_string_literal: true

require 'sinatra'

require_relative 'lib/controller/controller'
require_relative 'lib/services/video_games'
require_relative 'lib/storage/in_memory'

storage = InMemory.new
service = VideoGames.new storage
controller = Controller.new service

get '/' do
    'Hello World!'
end

get '/video_games' do
    controller.get_all_video_games
end