# frozen_string_literal: true

require_relative '../../lib/services/video_games'

describe VideoGames do
    let(:storage) { double 'mock storage' }
    subject { VideoGames.new storage }

    describe '#get_all' do
        it 'should get_all video games from storage and return them' do
            # Arrange
            video_games = []
            allow(storage).to receive(:get_all).and_return video_games

            # Act
            retrieved_video_games = subject.get_all

            # Assert
            expect(retrieved_video_games).to be video_games
        end
    end
end