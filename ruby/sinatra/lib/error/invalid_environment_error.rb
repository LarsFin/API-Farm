# frozen_string_literal: true

# Unique error message for when an environment doesn't exist
class InvalidEnvironmentError < StandardError
    def initialize(environment)
        super "Given environment #{environment} doesn't exist"
    end
end