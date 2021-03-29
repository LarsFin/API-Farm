# frozen_string_literal: true

require 'json'

# couples sinatra framework with services
class Controller
    def initialize(video_games_service)
        @video_games_service = video_games_service
    end

    def get_video_games
        video_games = @video_games_service.get_all

        [
          200,
          { 'Content-Type' => 'application/json' },
          video_games.to_json
        ]
    end

    def add_video_game request
        video_game_data = JSON.parse request.body.read
        addition = @video_games_service.add video_game_data

        if addition.fail_reason
            [
                400,
                { 'Content-Type' => 'text/plain' },
                addition.fail_reason
            ]
        else
            [
                201,
                { 'Content-Type' => 'application/json' },
                addition.result.to_json
            ]
        end
    end
end