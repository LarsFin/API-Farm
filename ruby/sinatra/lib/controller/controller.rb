# frozen_string_literal: true

# couples sinatra framework with services
class Controller
    def initialize video_games_service
        @video_games_service = video_games_service
    end

    def get_all_video_games
        video_games = @video_games_service.get_all
        [
            200,
            { 'Content-Type': 'application/json' },
            video_games
        ]
    end
end