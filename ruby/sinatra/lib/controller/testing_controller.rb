# frozen_string_literal: true

# handles api testing setup
class TestingController
    def initialize(data_loader, storage)
        @data_loader = data_loader
        @storage = storage
    end

    def setup; end
end