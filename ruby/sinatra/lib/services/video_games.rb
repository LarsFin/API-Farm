# frozen_string_literal: true

require 'date'

# service managing data from middleware, storage and the video game model
class VideoGames
    def initialize(storage, video_game_class)
        @storage = storage
        @video_game_class = video_game_class
    end

    def get_all
        @storage.get_all
    end

    def get(id)
        video_game = @storage.get id

        return { fail_code: 404, fail_reason: not_found_msg(id) } unless video_game

        { result: video_game }
    end

    def add(video_game_data)
        fail_result = validate_keys video_game_data

        return fail_result if fail_result

        video_game = @video_game_class.new

        fail_result = validate_required_properties video_game_data

        return fail_result if fail_result

        fail_result = set_properties video_game, video_game_data

        return fail_result if fail_result

        stored_video_game = @storage.add video_game

        { result: stored_video_game }
    end

    def update(id, video_game_data)
        fail_result = validate_keys video_game_data

        return fail_result if fail_result

        video_game = @storage.get id

        return { fail_code: 404, fail_reason: not_found_msg(id) } unless video_game

        fail_result = set_properties video_game, video_game_data

        return fail_result if fail_result

        updated_video_game = @storage.update id, video_game

        { result: updated_video_game }
    end

    def delete(id)
        video_game = @storage.delete id

        return { fail_code: 404, fail_reason: not_found_msg(id) } unless video_game

        { result: "Deleted video game with id '#{id}'." }
    end

  private

    def not_found_msg(id)
        "No video game with id '#{id}' could be found."
    end

    def validate_keys(video_game_data)
        video_game_data.each_key do |key|
            unless @video_game_class.method_defined? key.to_sym
              return { fail_code: 400,
                       fail_reason: "The provided data has an invalid attribute '#{key}'." }
            end
        end

        nil
    end

    def validate_required_properties(video_game_data)
        return { fail_reason: 'A name is required for a video game.' } unless video_game_data['name']

        return { fail_reason: 'A date_released is required for a video game.' } unless video_game_data['date_released']
    end

    def set_properties(video_game, video_game_data)
        set_name video_game, video_game_data
        set_developers video_game, video_game_data
        set_publishers video_game, video_game_data
        set_directors video_game, video_game_data
        set_producers video_game, video_game_data
        set_designers video_game, video_game_data
        set_programmers video_game, video_game_data
        set_artists video_game, video_game_data
        set_composers video_game, video_game_data
        set_platforms video_game, video_game_data

        fail_result = set_date video_game, video_game_data
        return fail_result if fail_result
    end

    def set_name(video_game, video_game_data)
        video_game.name = video_game_data['name']
    end

    def set_developers(video_game, video_game_data)
        video_game.developers = video_game_data['developers'] || []
    end

    def set_publishers(video_game, video_game_data)
        video_game.publishers = video_game_data['publishers'] || []
    end

    def set_directors(video_game, video_game_data)
        video_game.directors = video_game_data['directors'] || []
    end

    def set_producers(video_game, video_game_data)
        video_game.producers = video_game_data['producers'] || []
    end

    def set_designers(video_game, video_game_data)
        video_game.designers = video_game_data['designers'] || []
    end

    def set_programmers(video_game, video_game_data)
        video_game.programmers = video_game_data['programmers'] || []
    end

    def set_artists(video_game, video_game_data)
        video_game.artists = video_game_data['artists'] || []
    end

    def set_composers(video_game, video_game_data)
        video_game.composers = video_game_data['composers'] || []
    end

    def set_platforms(video_game, video_game_data)
        video_game.platforms = video_game_data['platforms'] || []
    end

    def set_date(video_game, video_game_data)
        return unless video_game_data['date_released']

        begin
            video_game.date_released = Date.parse video_game_data['date_released']
        rescue StandardError
            return { fail_code: 400,
                     fail_reason: "The provided date_released '#{video_game_data['date_released']}' is invalid." }
        end

        nil
    end
end