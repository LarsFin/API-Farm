# frozen_string_literal: true

require 'json'

# loads and transforms all data resources into in settings hash
class Config
    attr_reader :settings

    def initialize(environment)
        if environment == 'DEV'
            dev_file = File.read 'config.dev.json'

            @settings = JSON.parse dev_file
        else
            abort "Given environment doesn't exist"
        end
    end
end