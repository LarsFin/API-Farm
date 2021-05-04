# frozen_string_literal: true

require_relative '../../lib/models/video_game'

describe VideoGame do
    describe '#to_hash' do
        it 'should return object as hash' do
            # Arrange
            subject.id = 5
            subject.name = 'Dodo Alone'
            subject.developers = 'Astral Dawn Studios'
            subject.publishers = []
            subject.directors = []
            subject.producers = ['Lars Finlay']
            subject.designers = ['Karsten Finlay']
            subject.programmers = ['Lars Finlay']
            subject.artists = ['Jack Hopkins']
            subject.composers = ['Jack Hopkins']
            subject.platforms = ['Windows OS']

            date_released = double 'date released'
            subject.date_released = date_released
            date_released_s = '04/12/2002'
            allow(date_released).to receive(:strftime).with('%d/%m/%Y').and_return date_released_s

            # Act
            result = subject.to_hash

            # Assert
            expect(result).to be_instance_of Hash
            expect(result['id']).to be subject.id
            expect(result['name']).to be subject.name
            expect(result['developers']).to be subject.developers
            expect(result['publishers']).to be subject.publishers
            expect(result['directors']).to be subject.directors
            expect(result['producers']).to be subject.producers
            expect(result['designers']).to be subject.designers
            expect(result['programmers']).to be subject.programmers
            expect(result['artists']).to be subject.artists
            expect(result['composers']).to be subject.composers
            expect(result['platforms']).to be subject.platforms
            expect(result['date_released']).to be date_released_s
        end
    end
end