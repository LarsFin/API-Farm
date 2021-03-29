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
end