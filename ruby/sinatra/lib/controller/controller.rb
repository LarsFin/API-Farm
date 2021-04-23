# frozen_string_literal: true

require 'json'

# couples sinatra framework with services
class Controller
    def initialize(video_games_service)
        @video_games_service = video_games_service
    end

    def get_video_games
        video_games = @video_games_service.get_all
        hashed_video_games = video_games.map(&:to_hash)
        ok hashed_video_games
    end

    def get_video_game(request)
        id_retrieval = get_id request
        return bad_request id_retrieval[:fail_reason] if id_retrieval[:fail_reason]

        video_game_query = @video_games_service.get id_retrieval[:result]

        determine_response video_game_query
    end

    def add_video_game(request)
        body_retrieval = get_body request
        return bad_request body_retrieval[:fail_reason] if body_retrieval[:fail_reason]

        addition = @video_games_service.add body_retrieval[:result]

        if addition[:fail_reason]
            bad_request addition[:fail_reason]
        else
            created addition[:result].to_hash
        end
    end

    def update_video_game(request)
        id_retrieval = get_id request
        return bad_request id_retrieval[:fail_reason] if id_retrieval[:fail_reason]

        body_retrieval = get_body request
        return bad_request body_retrieval[:fail_reason] if body_retrieval[:fail_reason]

        update = @video_games_service.update id_retrieval[:result], body_retrieval[:result]

        determine_response update
    end

    def delete_video_game(request)
        id_retrieval = get_id request
        return bad_request id_retrieval[:fail_reason] if id_retrieval[:fail_reason]

        subtraction = @video_games_service.delete id_retrieval[:result]

        determine_response subtraction
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

    def determine_response(attempt)
        case attempt[:fail_code]
        when 400 then bad_request attempt[:fail_reason]
        when 404 then not_found attempt[:fail_reason]
        else
            if attempt[:result].is_a? String
                ok_text attempt[:result]
            else
                ok attempt[:result].to_hash
            end
        end
    end

    def ok(result)
        [
            200,
            { 'Content-Type' => 'application/json' },
            result.to_json
        ]
    end

    def ok_text(result)
        [
            200,
            { 'Content-Type' => 'TODO: CHANGE LATER!' },
            result
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