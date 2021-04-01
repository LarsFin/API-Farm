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
        video_game = @video_game_class.new

        if !video_game_data['name']
            return { :fail_reason => 'A name is required for a video game.' }
        else
            video_game.name = video_game_data['name']
        end

        if !video_game_data['date_released']
            return { :fail_reason => 'A date_released is required for a video game.' }
        else
            video_game.date_released = Date.parse video_game_data['date_released'] rescue 
            return { :fail_reason => "The provided date_released '#{video_game_data['date_released']}' is invalid." }
        end

        video_game.developers = video_game_data['developers']
        video_game.publishers = video_game_data['publishers']
        video_game.directors = video_game_data['directors']
        video_game.producers = video_game_data['producers']
        video_game.designers = video_game_data['designers']
        video_game.programmers = video_game_data['programmers']
        video_game.artists = video_game_data['artists']
        video_game.composers = video_game_data['composers']
        video_game.platforms = video_game_data['platforms']

        stored_video_game = @storage.add video_game

        return { :result => stored_video_game }
    end
end