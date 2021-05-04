# frozen_string_literal: true

# Unique error message for when an environment doesn't exist
class InvalidEnvironmentError < StandardError
    attr_reader :environment

    def initialize(environment)
        @environment = environment
        super "Given environment #{environment} doesn't exist"
    end
end