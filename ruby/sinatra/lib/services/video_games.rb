# frozen_string_literal: true

# service managing data from middleware, storage and the video game model
class VideoGames
    def initialize(storage)
        @storage = storage
    end
end