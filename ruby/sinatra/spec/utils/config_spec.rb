# frozen_string_literal: true

require_relative '../../lib/utils/config'

describe Config do
    environment = 'DEV'
    subject { Config.new environment }

    it 'should read and convert environment dev file' do
        # Arrange
        dev_file = double 'DEV file path'
        config_item = double 'Data file resource'
        config_array = [config_item]
        allow(File).to receive(:read).with('config.dev.json').and_return dev_file
        allow(JSON).to receive(:parse).with(dev_file).and_return config_array

        # Act
        result = subject.settings

        # Assert
        expect(result).to be config_array
    end
end

describe Config do
    environment = 'N/A'
    subject { Config.new environment }

    it 'should read and convert environment dev file' do
        # Arrange
        dev_file = double 'DEV file path'
        config_item = double 'Data file resource'
        config_array = [config_item]
        allow(File).to receive(:read).with('config.dev.json').and_return dev_file
        allow(JSON).to receive(:parse).with(dev_file).and_return config_array

        # Act
        result = subject.settings

        # Assert
        expect(result).to eq "Given environment doesn't exist"
    end
end