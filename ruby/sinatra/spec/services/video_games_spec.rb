# frozen_string_literal: true

require_relative '../../lib/services/video_games'

describe VideoGames do
    let(:storage) { double 'mock storage' }
    subject { VideoGames.new storage }

    describe '#get_all' do
        it 'should get_all video games from storage and return them' do
            # Arrange
            video_games = []
            video_games_json = 'video games as json'
            allow(storage).to receive(:get_all).and_return video_games
            allow(video_games).to receive(:to_json).and_return video_games_json

            # Act
            retrieved_video_games = subject.get_all

            # Assert
            expect(retrieved_video_games).to be video_games_json
        end
    end
end