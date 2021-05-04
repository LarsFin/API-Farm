# frozen_string_literal: true

require 'json'
require_relative '../error/invalid_environment_error'

# loads and transforms all data resources into in settings hash
class Config
    attr_reader :settings

    def initialize(environment)
        dup_environment = environment.dup
        dup_environment.upcase!

        case dup_environment
        when 'DEV'
            set_file 'config.dev.json'
        when 'PROD'
            set_file 'config.prod.json'
        when 'TEST'
            set_file 'config.test.json'
        else
            raise InvalidEnvironmentError, environment
        end
    end

  private

    def set_file(file_path)
        dev_file = File.read file_path

        @settings = JSON.parse dev_file
    end
end