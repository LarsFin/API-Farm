using ApiFarm.Controllers;
using ApiFarm.Models.Impl;
using ApiFarm.Repositories;
using ApiFarm.Test.Helpers;
using ApiFarm.Utils;
using Microsoft.AspNetCore.Http;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Controllers
{
    public class TestingControllerTests
    {
        private TestingController subject;

        private Mock<IDataLoader<VideoGame>> mockVideoGameDataLoader;
        private Mock<IRepository<VideoGame>> mockVideoGameStorage;

        [SetUp]
        public void SetUp()
        {
            mockVideoGameDataLoader = new Mock<IDataLoader<VideoGame>>();
            mockVideoGameStorage = new Mock<IRepository<VideoGame>>();

            subject = new TestingController(mockVideoGameDataLoader.Object, mockVideoGameStorage.Object);
        }

        private class SetUpTestsShould : TestingControllerTests
        {
            [Test]
            public void LoadAndSerializeDataThenAddToStorage()
            {
                // Arrange
                var sampleVideoGame1 = new VideoGame();
                var sampleVideoGame2 = new VideoGame();
                var sampleVideoGames = new VideoGame[]
                {
                    sampleVideoGame1,
                    sampleVideoGame2,
                };

                mockVideoGameDataLoader.Setup(m => m.Load(TestingController.DataSamplePath)).Returns(sampleVideoGames);

                // Act
                var objectResult = subject.SetUpTests().AsObjectResult();

                // Assert
                objectResult.StatusCode.ShouldBe(StatusCodes.Status200OK);
                objectResult.Value.ShouldBe(TestingController.SuccessMessage);

                mockVideoGameStorage.Verify(m => m.Add(sampleVideoGame1));
                mockVideoGameStorage.Verify(m => m.Add(sampleVideoGame2));
            }
        }
    }
}
