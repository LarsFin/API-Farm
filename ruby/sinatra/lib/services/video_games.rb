# frozen_string_literal: true

require 'Date'

# service managing data from middleware, storage and the video game model
class VideoGames
    def initialize(storage, video_game_class)
        @storage = storage
        @video_game_class = video_game_class
    end

    def get_all
        @storage.get_all
    end

    def add video_game_data
    end
end