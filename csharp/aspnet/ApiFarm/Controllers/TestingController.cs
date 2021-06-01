using ApiFarm.Models.Impl;
using ApiFarm.Repositories;
using ApiFarm.Utils;
using Microsoft.AspNetCore.Mvc;

namespace ApiFarm.Controllers
{
    /// <summary>
    /// Responsible for exposing an endpoint for setting up sample data for API testing.
    /// </summary>
    [ApiController]
    public class TestingController : ControllerBase
    {
        /// <summary>
        /// The path to the data file containing sample <see cref="VideoGame"/> objects for API tests.
        /// </summary>
        public const string DataSamplePath = "./data.json";

        /// <summary>
        /// Response message for when data was successfully loaded into storage for API testing.
        /// </summary>
        public const string SuccessMessage = "Successfully loaded data.";

        private IDataLoader<VideoGame> videoGameDataLoader;
        private IRepository<VideoGame> videoGameStorage;

        /// <summary>
        /// Initializes a new instance of the <see cref="TestingController"/> class.
        /// </summary>
        /// <param name="videoGameDataLoader">Dependency to load sample data and serialize into <see cref="VideoGame"/> entities.</param>
        /// <param name="videoGameStorage">Dependency for storgin <see cref="VideoGame"/> instances.</param>
        public TestingController(
            IDataLoader<VideoGame> videoGameDataLoader,
            IRepository<VideoGame> videoGameStorage)
        {
            this.videoGameDataLoader = videoGameDataLoader;
            this.videoGameStorage = videoGameStorage;
        }

        /// <summary>
        /// Loads sample data and adds serialized <see cref="VideoGame"/> instances to storage.
        /// </summary>
        /// <returns>Confirmation as to whether API tests have been successfully setup.</returns>
        [HttpGet]
        [Route("/api_tests/setup")]
        public ActionResult<string> SetUpTests()
        {
            var sampleVideoGames = this.videoGameDataLoader.Load(DataSamplePath);

            foreach (var videoGame in sampleVideoGames)
            {
                this.videoGameStorage.Add(videoGame);
            }

            return this.Ok(SuccessMessage);
        }
    }
}
