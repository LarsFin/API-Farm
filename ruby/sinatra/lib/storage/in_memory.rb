# frozen_string_literal: true

# in memory storage
class InMemory
    attr_accessor :video_games, :counter

    def initialize
        init
    end

    def add(video_game)
        @counter += 1
        video_game.id = @counter
        @video_games << video_game
        video_game
    end

    def get(id)
        @video_games.each do |video_game|
            return video_game if video_game.id == id
        end

        nil
    end

    def get_all
        @video_games
    end

    def update(id, video_game)
        i = 0

        while i < @video_games.length
            if @video_games[i].id == id
                @video_games[i] = video_game
                return @video_games[i]
            end

            i += 1
        end
    end

    def delete(id)
        i = 0

        while i < @video_games.length
            return @video_games.delete_at i if @video_games[i].id == id

            i += 1
        end
    end

    def reset
        init
    end

  private

    def init
        @video_games = []
        @counter = 0
    end
end