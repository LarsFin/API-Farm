# frozen_string_literal: true

require_relative '../../lib/storage/in_memory'

describe InMemory do
    describe '#add' do
        it 'should add video game to array' do
            # Arrange
            video_game = double 'mock video game'

            # Act
            subject.add video_game

            # Assert
            expect(subject.video_games.length).to eq 1
            expect(subject.video_games[0]).to be video_game
        end
    end
end