from flask.lib.storage.in_memory import InMemory
from unittest.mock import Mock

def test_should_add_game_to_array():
    # Arrange
    mock = Mock()
    video_game = mock
    video_games = []
    in_memory = InMemory(video_games)
    
    # Act
    in_memory.add(video_game)
    
    # Assert
    assert len(video_games) == 1
    assert in_memory.video_games == [video_game]

def test_should_get_video_game_from_array():
    # Arrange
    mock_1 = Mock(id = 1)
    mock_2 = Mock(id = 2)
    mock_3 = Mock(id = 3)
    video_game_1 = mock_1
    video_game_2 = mock_2
    video_game_3 = mock_3
    video_games = [video_game_1, video_game_2, video_game_3]
    in_memory = InMemory(video_games)

    # Act
    retrieved_video_game = in_memory.get(2)

    # Assert
    assert retrieved_video_game == video_game_2

def test_should_return_nil_when_no_video_game_has_passed_id():
    # Arrange
    mock = Mock(id = 1)
    video_game = mock

    # Act
    retrieved_video_game = in_memory.get

    # Assert
    assert retrieved_video_game == nil