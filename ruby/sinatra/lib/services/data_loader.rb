# frozen_string_literal: true

require 'json'

# loads and transforms all data resources into in memory objects
class DataLoader
    def initialize(data_file_path, video_game_class)
        @data_file_path = data_file_path
        @video_game_class = video_game_class
    end

    def load
        data_file_contents = File.read @data_file_path
        data_resources = JSON.parse data_file_contents

        data_resources.map do |resource|
            init_video_game resource
        end
    end

  private

    def init_video_game(video_game_hash)
        video_game = @video_game_class.new

        set_name video_game, video_game_hash
        set_developers video_game, video_game_hash
        set_publishers video_game, video_game_hash
        set_directors video_game, video_game_hash
        set_producers video_game, video_game_hash
        set_designers video_game, video_game_hash
        set_programmers video_game, video_game_hash
        set_artists video_game, video_game_hash
        set_composers video_game, video_game_hash
        set_platforms video_game, video_game_hash
        set_date_released video_game, video_game_hash

        video_game
    end

    def set_name(video_game, video_game_hash)
        video_game.name = video_game_hash['name']
    end

    def set_developers(video_game, video_game_hash)
        video_game.developers = video_game_hash['developers']
    end

    def set_publishers(video_game, video_game_hash)
        video_game.publishers = video_game_hash['publishers']
    end

    def set_directors(video_game, video_game_hash)
        video_game.directors = video_game_hash['directors']
    end

    def set_producers(video_game, video_game_hash)
        video_game.producers = video_game_hash['producers']
    end

    def set_designers(video_game, video_game_hash)
        video_game.designers = video_game_hash['designers']
    end

    def set_programmers(video_game, video_game_hash)
        video_game.programmers = video_game_hash['programmers']
    end

    def set_artists(video_game, video_game_hash)
        video_game.artists = video_game_hash['artists']
    end

    def set_composers(video_game, video_game_hash)
        video_game.composers = video_game_hash['composers']
    end

    def set_platforms(video_game, video_game_hash)
        video_game.platforms = video_game_hash['platforms']
    end

    def set_date_released(video_game, video_game_hash)
        video_game.date_released = Date.parse video_game_hash['date_released']
    end
end