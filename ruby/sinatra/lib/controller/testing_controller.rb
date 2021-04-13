# frozen_string_literal: true

# handles api testing setup
class TestingController
    def initialize(data_loader, storage)
        @data_loader = data_loader
        @storage = storage
    end

    def setup
        @storage.reset
        
        video_games = @data_loader.load

        for video_game in video_games
            @storage.add video_game
        end
    end
end