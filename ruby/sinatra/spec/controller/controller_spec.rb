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

    describe '#add_video_game' do
        it 'should return correct response with 201 status code' do
            # Arrange
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game as json'
            video_game_data = double 'video game data'
            addition = double 'attempt to add video game'
            result = double 'created video game result'
            json_result = double 'created video game result as json'

            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            JSON.stub(:parse).with(json_video_game).and_return video_game_data
            allow(video_games_service).to receive(:add).with(video_game_data).and_return addition
            allow(addition).to receive(:fail_reason)
            allow(addition).to receive(:result).and_return result
            allow(result).to receive(:to_json).and_return json_result

            # Act
            response = subject.add_video_game request

            # Assert
            expect(response[0]).to eq 201
            expected_headers = { 'Content-Type' => 'application/json' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be json_result
        end

        it 'should return correct response with 400 status code' do
            # Arrange
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game as json'
            video_game_data = double 'video game data'
            addition = double 'attempt to add video game'
            fail_reason = 'No name was provided'

            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            JSON.stub(:parse).with(json_video_game).and_return video_game_data
            allow(video_games_service).to receive(:add).with(video_game_data).and_return addition
            allow(addition).to receive(:fail_reason).and_return fail_reason

            # Act
            response = subject.add_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
        end
    end
end