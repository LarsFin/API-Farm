using Microsoft.AspNetCore.Mvc;

namespace ApiFarm.Controllers
{
    /// <summary>
    /// Controller for checking the health of the ApiFarm resource.
    /// </summary>
    [ApiController]
    public class HealthController : ControllerBase
    {
        /// <summary>
        /// Endpoint to check that the ApiFarm service is serving.
        /// </summary>
        /// <returns>pong, indicating service is ready.</returns>
        [HttpGet]
        [Route("/ping")]
        public string Ping() => "pong";
    }
}
