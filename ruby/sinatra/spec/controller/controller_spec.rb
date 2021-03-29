# frozen_string_literal: true

require_relative '../../lib/controller/controller'

describe Controller do
    let(:video_games_service) { double 'video game service' }
    subject { Controller.new video_games_service }

    describe '#get_video_games' do
        it 'should return correct response with 200 status code' do
            # Arrange
            video_games = double 'video games'
            json_video_games = double 'video games as json'
            allow(video_games_service).to receive(:get_all).and_return video_games
            allow(video_games).to receive(:to_json).and_return json_video_games

            # Act
            response = subject.get_video_games

            # Assert
            expect(response[0]).to eq 200
            expected_headers = { 'Content-Type' => 'application/json' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be json_video_games
        end
    end
end