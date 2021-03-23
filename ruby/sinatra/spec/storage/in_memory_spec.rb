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

        it 'should return video game with id which increments' do
            # Arrange
            video_game_1 = double 'first video game'
            video_game_2 = double 'second video game'

            # Act
            created_video_game_1 = subject.add video_game_1
            created_video_game_2 = subject.add video_game_2

            # Assert
            expect(video_game_1).to receive(:id=).with 1
            expect(created_video_game_1).to be video_game_1
            
            expect(video_game_2).to receive(:id=).with 2
            expect(created_video_game_2).to be video_game_2
        end
    end

    describe '#get' do
        it 'should get video game from array with passed index' do
            # Arrange
            subject.video_games << double('first video game', id: 1)
            subject.video_games << double('second video game', id: 2)
            subject.video_games << double('third video game', id: 3)

            # Act
            retrieved_video_game = subject.get 2

            # Assert
            expect(retrieved_video_game).to be subject.video_games[1]
        end

        it 'should return nil when no video game has passed id' do
            # Arrange
            subject.video_games << double('video game', id: 1)

            # Act
            retrieved_video_game = subject.get 2

            # Assert
            expect(retrieved_video_game).to eq nil
        end
    end

    describe '#get_all' do
        it 'should return all video games' do
            # Arrange
            subject.video_games << double('video game 1')
            subject.video_games << double('video game 2')
            subject.video_games << double('video game 3')

            # Act
            retrieved_video_games = subject.get_all

            # Assert
            expect(retrieved_video_games.length).to eq 3
            expect(retrieved_video_games).to be subject.video_games
        end
    end

    describe '#update' do
        it 'should overwrite and return video game with passed id' do
            # Arrange
            subject.video_games << double('original video game', id: 1)
            updating_video_game = double 'updating video game'

            # Act
            updated_video_game = subject.update 1, updating_video_game

            # Assert
            expect(updated_video_game).to be updated_video_game
            expect(subject.video_games[0]).to be updating_video_game
        end

        it 'should return nil when no video game has passed id' do
            # Arrange
            subject.video_games << double('video game', id: 1)
            updating_video_game = double 'updating video game'

            # Act
            updated_video_game = subject.update 2, updating_video_game

            # Assert
            expect(updated_video_game).to eq nil
        end
    end

    describe '#delete' do
        it 'should remove and return video game with passed id' do
            # Arrange
            subject.video_games[0] = video_game_to_delete = double 'video game to delete', id: 1

            # Act
            deleted_video_game = subject.delete 1

            # Assert
            expect(subject.video_games.length).to eq 0
            expect(deleted_video_game).to be video_game_to_delete
        end

        it 'should return nil when no video game has passed id' do
            # Arrange
            subject.video_games << double('video game', id: 1)

            # Act
            deleted_video_game = subject.delete 2

            # Assert
            expect(subject.video_games.length).to eq 1
            expect(deleted_video_game).to eq nil
        end
    end
end