# frozen_string_literal: true

require_relative '../../lib/controller/testing_controller'

describe TestingController do
    let(:data_loader) { double 'Data Loader' }
    let(:storage) { double 'Storage' }
    subject { TestingController.new data_loader, storage }

    describe '#setup' do
        it 'should reset storage and add video games from data loader' do
            # Arrange
            video_game1 = double 'video game 1'
            video_game2 = double 'video game 2'
            video_game3 = double 'video game 3'
            video_games = [video_game1, video_game2, video_game3]
            allow(data_loader).to receive(:load).and_return video_games

            # Assert
            expect(storage).to receive(:reset)
            expect(storage).to receive(:add).with video_game1
            expect(storage).to receive(:add).with video_game2
            expect(storage).to receive(:add).with video_game3

            # Act
            subject.setup
        end
    end
end