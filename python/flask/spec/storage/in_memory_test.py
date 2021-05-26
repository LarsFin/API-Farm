from flask.lib.storage.in_memory import add

def test_should_add_game_to_array():
    in_memory = InMemory()
    in_memory.add(video_game)
    assert in_memory.video_games == [video_game]
