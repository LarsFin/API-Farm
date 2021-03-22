# frozen_string_literal: true

# in memory storage
class InMemory
    attr_accessor :video_games

    def initialize
        @video_games = []
    end

    def add(video_game)
        video_games << video_game
    end

    def get(id)
        video_games.each do |video_game|
            return video_game if video_game.id == id
        end

        nil
    end

    def get_all
    end

    def update(id, video_game)
        i = 0

        while i < video_games.length
            return video_games[i] = video_game if video_games[i].id == id

            i += 1
        end
    end

    def delete(id)
        i = 0

        while i < video_games.length
            return video_games.delete_at i if video_games[i].id == id

            i += 1
        end
    end
end