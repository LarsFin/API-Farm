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

    def get_video_game(id)
        id = id.to_i
        
        if id == 0 
            bad_request('No video game with this id')
        else
            video_game = @video_games_service.get id

            if video_game[:fail_reason]
                bad_request video_game[:fail_reason]
            else
                ok video_game
            end
        end
    end

    def add_video_game(request)
        body_retrieval = get_body request
        return bad_request body_retrieval[:fail_reason] if body_retrieval[:fail_reason]

        addition = @video_games_service.add body_retrieval[:result]

        if addition[:fail_reason]
            bad_request addition[:fail_reason]
        else
            created addition[:result]
        end
    end

    def update_video_game(request)
        id_retrieval = get_id request
        return bad_request id_retrieval[:fail_reason] if id_retrieval[:fail_reason]

        body_retrieval = get_body request
        return bad_request body_retrieval[:fail_reason] if body_retrieval[:fail_reason]

        update = @video_games_service.update id_retrieval[:result], body_retrieval[:result]

        determine_update_response update
    end

  private

    def get_id(request)
        id_s = request.params['id']
        id_i = id_s.to_i

        return { fail_reason: "The provided id '#{id_s}' is invalid." } if id_i <= 0

        { result: id_i }
    end

    def get_body(request)
        body = JSON.parse request.body.read
        { result: body }
    rescue JSON::ParserError
        { fail_reason: 'Invalid JSON in body' }
    end

    def determine_update_response(update_attempt)
        case update_attempt[:fail_code]
        when 400 then bad_request update_attempt[:fail_reason]
        when 404 then not_found update_attempt[:fail_reason]
        else ok update_attempt[:result]
        end
    end

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

    def not_found(reason)
        [
            404,
            { 'Content-Type' => 'text/plain' },
            reason
        ]
    end
end