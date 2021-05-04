# frozen_string_literal: true

require 'json'
require_relative '../error/invalid_environment_error'

# loads and transforms all data resources into in settings hash
class Config
    attr_reader :settings

    def initialize(environment)
        raise InvalidEnvironmentError, environment unless %w[DEV TEST PROD].include? environment.upcase

        set_file "config.#{environment.downcase}.json"
    end

  private

    def set_file(file_path)
        dev_file = File.read file_path

        @settings = JSON.parse dev_file
    end
end