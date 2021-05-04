# frozen_string_literal: true

require 'json'

# loads and transforms all data resources into in settings hash
class Config
    attr_reader :settings

    def initialize(environment)
        case environment.upcase
        when 'DEV'
            set_file 'config.dev.json'
        when 'PROD'
            set_file 'config.prod.json'
        when 'TEST'
            set_file 'config.test.json'
        else
            raise "Given environment doesn't exist"
        end
    end

  private

    def set_file(file_path)
        dev_file = File.read file_path

        @settings = JSON.parse dev_file
    end
end