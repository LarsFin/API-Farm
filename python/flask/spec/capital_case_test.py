def capital_case(_x):
    return _x.capitalize()

def test_capital_case():
    assert capital_case('semaphore') == 'Semaphore'
