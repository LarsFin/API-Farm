class InvalidEnvironmentError < StandardError
    attr_reader :environment

    def initialize(environment)
        @environment = environment
        super "Given environment #{environment} doesn't exist"
    end
end