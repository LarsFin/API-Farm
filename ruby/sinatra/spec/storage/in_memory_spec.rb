# frozen_string_literal: true

require_relative '../../lib/storage/in_memory'

describe InMemory do
    describe '#add' do
        it 'should add video game to array' do
            # Arrange
            video_game = double 'mock video game'
            allow(video_game).to receive(:to_hash)

            # Assert
            expect(video_game).to receive(:id=).with 1

            # Act
            subject.add video_game

            # Assert
            expect(subject.video_games.length).to eq 1
            expect(subject.video_games[0]).to be video_game
        end

        it 'should return video game with id which increments' do
            # Arrange
            video_game1 = double 'first video game'
            video_game1_hash = double 'first video game as hash'
            video_game2 = double 'second video game'
            video_game2_hash = double 'second video game as hash'
            allow(video_game1).to receive(:to_hash).and_return video_game1_hash
            allow(video_game2).to receive(:to_hash).and_return video_game2_hash

            # Assert
            expect(video_game1).to receive(:id=).with 1
            expect(video_game2).to receive(:id=).with 2

            # Act
            created_video_game1 = subject.add video_game1
            created_video_game2 = subject.add video_game2

            # Assert
            expect(created_video_game1).to be video_game1_hash
            expect(created_video_game2).to be video_game2_hash
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
            video_game1 = double 'video game 1'
            video_game1_hash = double 'video game 1 as hash'
            video_game2 = double 'video game 2'
            video_game2_hash = double 'video game 2 as hash'
            video_game3 = double 'video game 3'
            video_game3_hash = double 'video game 3 as hash'
            allow(video_game1).to receive(:to_hash).and_return video_game1_hash
            allow(video_game2).to receive(:to_hash).and_return video_game2_hash
            allow(video_game3).to receive(:to_hash).and_return video_game3_hash
            subject.video_games << video_game1 << video_game2 << video_game3
            expected = [
                video_game1_hash,
                video_game2_hash,
                video_game3_hash
            ]

            # Act
            retrieved_video_games = subject.get_all

            # Assert
            expect(retrieved_video_games.length).to eq 3
            expect(retrieved_video_games).to eq expected
        end
    end

    describe '#update' do
        it 'should overwrite and return hashed video game with passed id' do
            # Arrange
            updating_video_game = double 'video game used to update'
            subject.video_games << double('video game to pass', id: 1)
            subject.video_games << double('video game to update', id: 2)
            updated_video_game_hash = double 'updated video game as hash'
            allow(updating_video_game).to receive(:to_hash).and_return updated_video_game_hash

            # Act
            updated_video_game = subject.update 2, updating_video_game

            # Assert
            expect(updated_video_game).to be updated_video_game_hash
            expect(subject.video_games[1]).to be updating_video_game
        end

        it 'should return nil when no video game has passed id' do
            # Arrange
            subject.video_games << double('video game', id: 1)
            updating_video_game = double 'video game used to update'

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

    describe '#reset' do
        it 'should reset store array and counter' do
            # Arrange
            subject.video_games = [1, 2, 3, 4, 5]
            subject.counter = 99

            # Act
            subject.reset

            # Assert
            expect(subject.video_games).to be_empty
            expect(subject.counter).to be 0
        end
    end
end