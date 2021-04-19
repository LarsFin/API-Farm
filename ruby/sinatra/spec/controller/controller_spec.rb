# frozen_string_literal: true

require_relative '../../lib/controller/controller'

describe Controller do
    let(:video_games_service) { double 'video game service' }
    subject { Controller.new video_games_service }

    describe '#get_video_games' do
        it 'should return correct response with 200 status code' do
            # Arrange
            video_games = double 'video games'
            hash_video_games = double 'video games as hash'
            json_video_games = double 'video games as json'
            allow(video_games_service).to receive(:get_all).and_return video_games
            allow(video_games).to receive(:map).and_return hash_video_games
            allow(hash_video_games).to receive(:to_json).and_return json_video_games

            # Act
            response = subject.get_video_games

            # Assert
            expect(response[0]).to eq 200
            expected_headers = { 'Content-Type' => 'application/json' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be json_video_games
        end
    end

    describe '#get_video_game' do
        it 'Should return a specific video game via id given' do
            # Arrange
            id = '1'
            request_params = { 'id' => id }
            request = double 'request'
            video_game_query = double 'attempt to get video game'
            result = double 'video game obtained'
            hash_result = double 'video game obtained as a hash'
            json_result = double 'video game obtained as a json'

            allow(request).to receive(:params).and_return request_params
            allow(video_games_service).to receive(:get).with(id.to_i).and_return video_game_query
            allow(video_game_query).to receive(:[]).with :fail_code
            allow(video_game_query).to receive(:[]).with(:result).and_return result
            allow(result).to receive(:to_hash).and_return hash_result
            allow(hash_result).to receive(:to_json).and_return json_result

            # Act
            response = subject.get_video_game request

            # Assert
            expect(response[0]).to eq 200
            expected_headers = { 'Content-Type' => 'application/json' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be json_result
        end

        it 'should return correct response with 400 status code' do
            # Arrange
            id = 'not a valid id!'
            request_params = { 'id' => id }
            request = double 'request'

            allow(request).to receive(:params).and_return request_params

            # Act
            response = subject.get_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to eq "The provided id '#{id}' is invalid."
        end

        it 'should return correct response with 404 status code' do
            # Arrange
            id = '12'
            request_params = { 'id' => id }
            request = double 'request'
            video_game_query = double 'attempt to get video game'
            fail_reason = 'could not find video game with this id!'

            allow(request).to receive(:params).and_return request_params
            allow(video_games_service).to receive(:get).with(id.to_i).and_return video_game_query
            allow(video_game_query).to receive(:[]).with(:fail_code).and_return 404
            allow(video_game_query).to receive(:[]).with(:fail_reason).and_return fail_reason

            # Act
            response = subject.get_video_game request

            # Assert
            expect(response[0]).to eq 404
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
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
            hash_result = double 'created video game result as hash'
            json_result = double 'created video game result as json'

            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            allow(JSON).to receive(:parse).with(json_video_game).and_return video_game_data
            allow(video_games_service).to receive(:add).with(video_game_data).and_return addition
            allow(addition).to receive(:[]).with(:fail_reason)
            allow(addition).to receive(:[]).with(:result).and_return result
            allow(result).to receive(:to_hash).and_return hash_result
            allow(hash_result).to receive(:to_json).and_return json_result

            # Act
            response = subject.add_video_game request

            # Assert
            expect(response[0]).to eq 201
            expected_headers = { 'Content-Type' => 'application/json' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be json_result
        end

        it 'should return correct response with 400 status code from invalid JSON' do
            # Arrange
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game as json'
            fail_reason = 'Invalid JSON in body'

            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            allow(JSON).to receive(:parse).with(json_video_game).and_raise JSON::ParserError.new fail_reason

            # Act
            response = subject.add_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
        end

        it 'should return correct response with 400 status code from incorrect data' do
            # Arrange
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game as json'
            video_game_data = double 'video game data'
            addition = double 'attempt to add video game'
            fail_reason = 'No name was provided'

            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            allow(JSON).to receive(:parse).with(json_video_game).and_return video_game_data
            allow(video_games_service).to receive(:add).with(video_game_data).and_return addition
            allow(addition).to receive(:[]).with(:fail_reason).and_return fail_reason

            # Act
            response = subject.add_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
        end
    end

    describe '#update_video_game' do
        it 'should return correct response with 200 status code' do
            # Arrange
            id = '12'
            request_params = { 'id' => id }
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game data as json'
            video_game_data = double 'video game data'
            update = double 'attempt to update video game'
            result = double 'updated video game'
            hash_result = double 'updated video game result as hash'
            json_result = double 'updated video game result as json'

            allow(request).to receive(:params).and_return request_params
            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            allow(JSON).to receive(:parse).with(json_video_game).and_return video_game_data
            allow(video_games_service).to receive(:update).with(id.to_i, video_game_data).and_return update
            allow(update).to receive(:[]).with(:fail_code)
            allow(update).to receive(:[]).with(:result).and_return result
            allow(result).to receive(:to_hash).and_return hash_result
            allow(hash_result).to receive(:to_json).and_return json_result

            # Act
            response = subject.update_video_game request

            # Assert
            expect(response[0]).to eq 200
            expected_headers = { 'Content-Type' => 'application/json' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be json_result
        end

        it 'should return correct response with 400 status code from invalid id' do
            # Arrange
            id = 'not a valid id!'
            request_params = { 'id' => id }
            request = double 'request'

            allow(request).to receive(:params).and_return request_params

            # Act
            response = subject.update_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to eq "The provided id '#{id}' is invalid."
        end

        it 'should return correct response with 400 status code from invalid JSON' do
            # Arrange
            id = '12'
            request_params = { 'id' => id }
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game as json'
            fail_reason = 'Invalid JSON in body'

            allow(request).to receive(:params).and_return request_params
            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            allow(JSON).to receive(:parse).with(json_video_game).and_raise JSON::ParserError.new fail_reason

            # Act
            response = subject.update_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
        end

        it 'should return correct response with 400 status code from invalid data' do
            # Arrange
            id = '12'
            request_params = { 'id' => id }
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game data as json'
            video_game_data = double 'video game data'
            update = double 'attempt to update video game'
            fail_reason = double 'invalid json data!'

            allow(request).to receive(:params).and_return request_params
            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            allow(JSON).to receive(:parse).with(json_video_game).and_return video_game_data
            allow(video_games_service).to receive(:update).with(id.to_i, video_game_data).and_return update
            allow(update).to receive(:[]).with(:fail_code).and_return 400
            allow(update).to receive(:[]).with(:fail_reason).and_return fail_reason

            # Act
            response = subject.update_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
        end

        it 'should return correct response with 404 status code' do
            # Arrange
            id = '12'
            request_params = { 'id' => id }
            request = double 'request'
            request_body = double 'request body'
            json_video_game = double 'video game data as json'
            video_game_data = double 'video game data'
            update = double 'attempt to update video game'
            fail_reason = double 'could not find video game to update!'

            allow(request).to receive(:params).and_return request_params
            allow(request).to receive(:body).and_return request_body
            allow(request_body).to receive(:read).and_return json_video_game
            allow(JSON).to receive(:parse).with(json_video_game).and_return video_game_data
            allow(video_games_service).to receive(:update).with(id.to_i, video_game_data).and_return update
            allow(update).to receive(:[]).with(:fail_code).and_return 404
            allow(update).to receive(:[]).with(:fail_reason).and_return fail_reason

            # Act
            response = subject.update_video_game request

            # Assert
            expect(response[0]).to eq 404
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
        end
    end

    describe '#delete_video_game' do
        it 'should return correct response with 200 status code' do
            # Arrange
            id = '1'
            request_params = { 'id' => id }
            request = double 'request'
            subtraction = double 'attempt to remove video game'
            result = double 'deleted video game'

            allow(request).to receive(:params).and_return request_params
            allow(video_games_service).to receive(:delete).with(id.to_i).and_return subtraction
            allow(subtraction).to receive(:[]).with :fail_code
            allow(subtraction).to receive(:[]).with(:result).and_return result

            # Act
            response = subject.delete_video_game request

            # Assert
            expect(response[0]).to eq 200
            expected_headers = { 'Content-Type' => 'application/json' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be result
        end

        it 'should return correct response with 400 status code' do
            # Arrange
            id = 'not a valid id!'
            request_params = { 'id' => id }
            request = double 'request'

            allow(request).to receive(:params).and_return request_params

            # Act
            response = subject.delete_video_game request

            # Assert
            expect(response[0]).to eq 400
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to eq "The provided id '#{id}' is invalid."
        end

        it 'should return correct response with 404 status code' do
            # Arrange
            id = '12'
            request_params = { 'id' => id }
            request = double 'request'
            subtraction = double 'attempt to delete video game'
            fail_reason = 'could not find video game with this id!'

            allow(request).to receive(:params).and_return request_params
            allow(video_games_service).to receive(:delete).with(id.to_i).and_return subtraction
            allow(subtraction).to receive(:[]).with(:fail_code).and_return 404
            allow(subtraction).to receive(:[]).with(:fail_reason).and_return fail_reason

            # Act
            response = subject.delete_video_game request

            # Assert
            expect(response[0]).to eq 404
            expected_headers = { 'Content-Type' => 'text/plain' }
            expect(response[1]).to eq expected_headers
            expect(response[2]).to be fail_reason
        end
    end
end
