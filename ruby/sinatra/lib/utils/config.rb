# frozen_string_literal: true

require 'json'

# loads and transforms all data resources into in settings hash
class Config
    attr_reader :settings

    def initialize(environment)
        case environment
            when "DEV"
                dev_file = File.read 'config.dev.json'

                @settings = JSON.parse dev_file
            when "PROD"
                dev_file = File.read 'config.prod.json'

                @settings = JSON.parse dev_file
            when "TEST"
                dev_file = File.read 'config.test.json'

                @settings = JSON.parse dev_file
            else
                abort "Given environment doesn't exist"
        end
    end
end