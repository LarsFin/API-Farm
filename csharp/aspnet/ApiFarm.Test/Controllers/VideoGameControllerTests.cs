using System.Collections.Generic;
using ApiFarm.Controllers;
using ApiFarm.Models.Impl;
using ApiFarm.Services;
using ApiFarm.Utils;
using Microsoft.AspNetCore.Mvc;
using Moq;
using NUnit.Framework;
using Shouldly;

namespace ApiFarm.Test.Controllers
{
    public class VideoGameControllerTests
    {
        private VideoGamesController subject;

        private Mock<IService<VideoGame>> mockVideoGameService;

        [SetUp]
        protected void SetUp()
        {
            mockVideoGameService = new Mock<IService<VideoGame>>();

            subject = new VideoGamesController(mockVideoGameService.Object);
        }

        private class GetAllShould : VideoGameControllerTests
        {
            [Test]
            public void Return200Response()
            {
                // Arrange
                var storedVideoGames = new Mock<IEnumerable<VideoGame>>();
                var mockQuery = new Mock<IQuery<IEnumerable<VideoGame>>>();

                mockVideoGameService.Setup(m => m.GetAll()).Returns(mockQuery.Object);
                mockQuery.Setup(m => m.Result).Returns(storedVideoGames.Object);

                // Act
                var actionResult = subject.GetAll();

                // Assert
                actionResult.ShouldBeOfType<OkObjectResult>();
                actionResult.Value.ShouldBe(storedVideoGames.Object);
            }
        }
    }
}
