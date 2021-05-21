using ApiFarm.Controllers;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Controllers
{
    public class VideoGameControllerTests
    {
        private VideoGamesController subject;

        [SetUp]
        protected void SetUp()
        {
            subject = new VideoGamesController();
        }

        private class GetAllShould : VideoGameControllerTests
        {
            // Temporary test.
            [Test]
            public void ReturnHelloWorld()
            {
                // Act
                var result = subject.GetAll();

                // Assert
                result.ShouldBe("Hello World!");
            }
        }
    }
}
