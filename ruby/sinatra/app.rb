# frozen_string_literal: true

require 'sinatra'

require_relative 'lib/services/video_games'
require_relative 'lib/storage/in_memory'

storage = InMemory.new
service = VideoGames.new storage

get '/' do
    'Hello World!'
end

get '/video_games' do
    service.get_all
end