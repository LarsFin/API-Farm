from flask.lib.storage.in_memory import InMemory
from unittest.mock import Mock

def test_should_add_game_to_array():
    print("Are we here?")
    # Arrange
    mock = Mock()
    video_game = mock
    video_games = []
    in_memory = InMemory(video_games)
    
    # Act
    in_memory.add(video_game)
    
    # Assert
    print("We are in assert")
    print(len(video_games))
    assert len(video_games) == 1
    assert in_memory.video_games == [video_game]
