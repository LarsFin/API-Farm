# frozen_string_literal: true

require 'json'

# service managing data from middleware, storage and the video game model
class VideoGames
    def initialize(storage)
        @storage = storage
    end

    def get_all
        @storage.get_all.to_json
    end
end