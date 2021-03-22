# frozen_string_literal: true

require_relative '../../lib/services/video_games'

describe VideoGames do
    let(:storage) { double 'mock storage' }
    subject { VideoGames.new storage }
end