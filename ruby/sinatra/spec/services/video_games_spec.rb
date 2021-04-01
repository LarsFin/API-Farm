# frozen_string_literal: true

require_relative '../../lib/services/video_games'

describe VideoGames do
    let(:storage) { double 'mock storage' }
    let(:video_game_class) { double 'video game class' }
    subject { VideoGames.new storage, video_game_class }

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

    describe '#add' do
        it 'should create and add video game to storage and return stored instance' do
            # Arrange
            video_game_data = {
                'name' => 'Video Game II',
                'developers' => [ 'developer 1', 'developer 2' ],
                'publishers' => [ 'publisher 1', 'publisher 2' ],
                'directors' => [ 'director 1' ],
                'producers' => [ 'producer 1', 'producer 2' ],
                'designers' => [ 'designer 1' ],
                'programmers' => [ 'programmer 1', 'programmer 2' ],
                'artists' => [ 'artist 1', 'artist 2' ],
                'composers' => [ 'composer 1' ],
                'platforms' => [ 'platform 1', 'platform 2', 'platform 3' ],
                'date_released' => '18/02/2002'
            }
            converted_date_released = double 'date released as date object'
            video_game = double 'video game'
            stored_video_game = double 'video game added to storage'
            allow(video_game_class).to receive(:new).and_return video_game
            allow(Date).to receive(:parse).with(video_game_data['date_released']).and_return converted_date_released
            allow(storage).to receive(:add).with(video_game).and_return stored_video_game

            # Assert
            expect(video_game).to receive(:name=).with video_game_data['name']
            expect(video_game).to receive(:date_released=).with converted_date_released
            expect(video_game).to receive(:developers=).with video_game_data['developers']
            expect(video_game).to receive(:publishers=).with video_game_data['publishers']
            expect(video_game).to receive(:directors=).with video_game_data['directors']
            expect(video_game).to receive(:producers=).with video_game_data['producers']
            expect(video_game).to receive(:designers=).with video_game_data['designers']
            expect(video_game).to receive(:programmers=).with video_game_data['programmers']
            expect(video_game).to receive(:artists=).with video_game_data['artists']
            expect(video_game).to receive(:composers=).with video_game_data['composers']
            expect(video_game).to receive(:platforms=).with video_game_data['platforms']

            # Act
            addition = subject.add video_game_data

            # Assert
            expect(addition[:result]).to be stored_video_game
        end

        it 'should return failure when video game data has no name' do
            # Arrange
            video_game_data = { }
            converted_date_released = double 'date released as date object'
            video_game = double 'video game'
            allow(video_game_class).to receive(:new).and_return video_game

            # Act
            addition = subject.add video_game_data

            # Assert
            expect(addition[:fail_reason]).to eq 'A name is required for a video game.'
        end

        it 'should return failure when video game data has no date' do
            # Arrange
            video_game_data = {
                'name' => 'Video Game II'
            }
            converted_date_released = double 'date released as date object'
            video_game = double 'video game'
            allow(video_game_class).to receive(:new).and_return video_game

            # Assert
            expect(video_game).to receive(:name=).with video_game_data['name']

            # Act
            addition = subject.add video_game_data

            # Assert
            expect(addition[:fail_reason]).to eq 'A date_released is required for a video game.'
        end

        it 'should return failure when video game data has an invalid date' do
            # Arrange
            video_game_data = {
                'name' => 'Video Game II',
                'date_released' => '08/23/2007'
            }
            converted_date_released = double 'date released as date object'
            video_game = double 'video game'
            allow(video_game_class).to receive(:new).and_return video_game
            allow(Date).to receive(:parse).and_raise('Failed!')

            # Assert
            expect(video_game).to receive(:name=).with video_game_data['name']

            # Act
            addition = subject.add video_game_data

            # Assert
            expect(addition[:fail_reason]).to eq "The provided date_released '#{video_game_data['date_released']}' is invalid."
        end
    end
end