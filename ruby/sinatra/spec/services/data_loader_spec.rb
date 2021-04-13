require_relative '../../lib/services/data_loader'

describe DataLoader do
    data_file_path = 'path/to/data/file'
    let(:video_game_class) { double 'video game class' }
    subject { DataLoader.new data_file_path, video_game_class }

    it 'should read and convert data file resources' do
        # Arrange
        contents_string = 'Data file contents'
        contents_array = double 'Data file contents as array'
        contents_item = double 'Data file resource'
        video_game = double 'Created video game instance'
        allow(File).to receive(:read).with(data_file_path).and_return contents_string
        allow(JSON).to receive(:parse).with(contents_string).and_return contents_array
        allow(contents_array).to receive(:map).and_yield(contents_item)
                                              .and_yield(contents_item)
                                              .and_yield contents_item
        allow(video_game_class).to receive(:new).and_return video_game
        
        vg_name = "name of video game"
        vg_developers = "developers of video game"
        vg_publishers = "publishers of video game"
        vg_directors = "directors of video game"
        vg_producers = "producers of video game"
        vg_designers = "designers of video game"
        vg_programmers = "programmers of video game"
        vg_artists = "artists of video game"
        vg_composers = "composers of video game"
        vg_platforms = "platforms of video game"
        vg_date_released = "date released of video game"

        allow(contents_item).to receive(:[]).with('name').and_return vg_name
        allow(contents_item).to receive(:[]).with('developers').and_return vg_developers
        allow(contents_item).to receive(:[]).with('publishers').and_return vg_publishers
        allow(contents_item).to receive(:[]).with('directors').and_return vg_directors
        allow(contents_item).to receive(:[]).with('producers').and_return vg_producers
        allow(contents_item).to receive(:[]).with('designers').and_return vg_designers
        allow(contents_item).to receive(:[]).with('programmers').and_return vg_programmers
        allow(contents_item).to receive(:[]).with('artists').and_return vg_artists
        allow(contents_item).to receive(:[]).with('composers').and_return vg_composers
        allow(contents_item).to receive(:[]).with('platforms').and_return vg_platforms
        allow(contents_item).to receive(:[]).with('date_released').and_return vg_date_released

        # Assert
        expect(video_game).to receive(:name=).with(vg_name).exactly(3).times
        expect(video_game).to receive(:developers=).with(vg_developers).exactly(3).times
        expect(video_game).to receive(:publishers=).with(vg_publishers).exactly(3).times
        expect(video_game).to receive(:directors=).with(vg_directors).exactly(3).times
        expect(video_game).to receive(:producers=).with(vg_producers).exactly(3).times
        expect(video_game).to receive(:designers=).with(vg_designers).exactly(3).times
        expect(video_game).to receive(:programmers=).with(vg_programmers).exactly(3).times
        expect(video_game).to receive(:artists=).with(vg_artists).exactly(3).times
        expect(video_game).to receive(:composers=).with(vg_composers).exactly(3).times
        expect(video_game).to receive(:platforms=).with(vg_platforms).exactly(3).times
        expect(video_game).to receive(:date_released=).with(vg_date_released).exactly(3).times

        # Act
        result = subject.load

        # Assert
        expect(result.length).to eq 3
        expect(result[0]).to be video_game
        expect(result[1]).to be video_game
        expect(result[2]).to be video_game
    end
end