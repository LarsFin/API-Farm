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

    describe '#get' do
        it 'should get video game from array with passed index' do
            # Arrange
            subject.video_games[0] = double('first video game', id: 1)
            subject.video_games[1] = double('second video game', id: 2)
            subject.video_games[2] = double('second video game', id: 3)

            # Act
            retrieved_video_game = subject.get 2

            # Assert
            expect(retrieved_video_game).to be subject.video_games[1]
        end
    end
end