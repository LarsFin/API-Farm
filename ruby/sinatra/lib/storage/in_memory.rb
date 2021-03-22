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

    def get(index)
        video_games.each do |video_game|
            return video_game if video_game.id == index
        end

        nil
    end

    def update(index, video_game)
    end
end