# frozen_string_literal: true

require 'json'

# couples sinatra framework with services
class Controller
    def initialize(video_games_service)
        @video_games_service = video_games_service
    end

    def get_video_games
        video_games = @video_games_service.get_all
        ok video_games
    end

    def add_video_game(request)
        begin video_game_data = JSON.parse request.body.read
        rescue JSON::ParserError
            return bad_request 'Invalid JSON in body'
        end

        addition = @video_games_service.add video_game_data

        if addition[:fail_reason]
            bad_request addition[:fail_reason]
        else
            created addition[:result]
        end
    end

    def update_video_game(request)
    end

  private

    def ok(result)
        [
            200,
            { 'Content-Type' => 'application/json' },
            result.to_json
        ]
    end

    def created(result)
        [
            201,
            { 'Content-Type' => 'application/json' },
            result.to_json
        ]
    end

    def bad_request(reason)
        [
            400,
            { 'Content-Type' => 'text/plain' },
            reason
        ]
    end
end